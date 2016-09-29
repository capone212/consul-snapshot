#!/bin/bash
#make build-all
# Zip and copy to the dist dir
echo "==> Packaging..."
for PLATFORM in $(find ./pkg -mindepth 1 -maxdepth 1 -type d); do
	OSARCH=$(basename ${PLATFORM})
        echo "--> ${OSARCH}"
	pushd $PLATFORM >/dev/null 2>&1
	zip ../${OSARCH}.zip ./*
        popd >/dev/null 2>&1
done
