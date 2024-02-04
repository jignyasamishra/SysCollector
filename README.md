# SysCollector
The SysCollector is a command-line utility written in Go that provides real-time information about CPU, memory, and disk usage on your system. It utilizes the shirou/gopsutil and spf13/cobra libraries to gather system statistics and manage command-line interactions.
# Features
- CPU Usage: Displays the current CPU usage as a percentage.
- Memory Usage: Provides information about total memory, free memory, and memory usage percentage.
- Disk Usage: Shows details on total disk space, free disk space, and disk usage percentage for the root directory.
# Flags
- -e, --export: Export system stats to a file. Specify the file path using this flag.
