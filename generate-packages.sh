#!/bin/bash

template=vec2
exclude=n.go

for pkg in vec4 vec8 vec16 vec32 vec64; do
  find "$pkg" -type f -not -name "$exclude" -exec rm "{}" ";"
  for file in $(basename -a $template/*); do
    if [ $file != $exclude ]; then
      sed "s/package *vec2/package ${pkg}/" "$template/$file" > ${pkg}/$file
    fi
  done
done

