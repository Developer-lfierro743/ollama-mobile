package discover

import (
    "os"
    "strings"
    "log/slog"
)

// DetectAndroidGPU detects Adreno/Mali GPUs on Android via /proc and sysfs
// since lspci is not available on Android/Bionic
func DetectAndroidGPU() (name string, vendor string, found bool) {
    // Check Adreno via Turnip/Mesa driver sysfs
    paths := []string{
        "/sys/class/kgsl/kgsl-3d0/gpu_model",
        "/sys/class/kgsl/kgsl-3d0/devfreq/cur_freq",
        "/proc/device-tree/model",
    }
    for _, path := range paths {
        if data, err := os.ReadFile(path); err == nil {
            content := strings.TrimSpace(string(data))
            if strings.Contains(strings.ToLower(content), "adreno") {
                slog.Info("Detected Adreno GPU", "source", path, "info", content)
                return content, "Qualcomm", true
            }
            if strings.Contains(strings.ToLower(content), "mali") {
                slog.Info("Detected Mali GPU", "source", path, "info", content)
                return content, "ARM", true
            }
        }
    }
    return "", "", false
}

// IsAndroidChroot detects if we're running inside an Android chroot
func IsAndroidChroot() bool {
    _, err := os.Stat("/proc/device-tree")
    return err == nil
}
