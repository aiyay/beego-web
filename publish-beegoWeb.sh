#!/bin/bash
# Program: system_init_shell
# Author Kevin
cd /Users/iamwlb/Documents/Workspace/go/src/beegoWeb
echo ' ' | sudo -S git add .
echo ' ' | sudo -S git commit -m "update"
echo ' ' | sudo -S git push origin master
echo "Complate"
echo "###############################################################"