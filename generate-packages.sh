#!/bin/bash

exclude=n.go
templateSize=16
for otherSize in 2 4 8 32 64; do
  pkg=vec$otherSize
  find "$pkg" -type f -not -name "$exclude" -exec rm "{}" ";"
  for file in $(basename -a $template/*); do
    if [ $file != $exclude ]; then
      sed "s/package *vec$templateSize/package ${pkg}/" "vec$templateSize/$file" | sed "s/Package *vec$templateSize/Package ${pkg}/" > ${pkg}/$file
    fi
  done
done

