# ollama-mobile

A mobile-optimized LLM runtime for Android with Vulkan/Adreno GPU support.

## Fork of [Ollama](https://github.com/ollama/ollama) (MIT License)

## What's Different
- Stripped desktop/GUI bloat (app/, harmony/, Docker)
- Vulkan backend patched for Android/Adreno GPU support
- Designed for rooted Android via chroot environment
- Tested on Adreno 710 (Snapdragon 6 Gen 1) with Turnip Mesa 26.1/Vulkan 1.4

## Requirements
- Rooted Android device
- Vulkan 1.1+ capable GPU (Adreno, Mali)
- Ubuntu chroot (via chroot-distro)
- BusyBox v1.36.1+

## Build
```bash
cmake -B build -DGGML_VULKAN=1
cmake --build build -j6
Credits
Original Ollama project: https://github.com/ollama/ollama
Turnip Mesa driver for Adreno GPU support
