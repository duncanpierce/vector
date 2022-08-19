#!/bin/bash

template=vec2

for file in $(basename -a $template/*); do
  if [ $file != 'n.go' ]; then
    for pkg in vec4 vec8 vec16 vec32 vec64; do
      sed "s/package *vec2/package ${pkg}/" "$template/$file" > ${pkg}/$file
    done
  fi
done

