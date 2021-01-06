#!/bin/bash
# Leer rutas
path=$(pwd)
bashrc=/etc/bash.bashrc
# Agregar ruta al PATH
echo $PATH | grep -q "(^\|:\")$path\(:\|/\{0,1\}$\)" || echo "PATH=\$PATH:$path" >> "$bashrc"; source "$bashrc"