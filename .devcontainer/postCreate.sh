#!/bin/bash
set -e

# Install Zig Language Server
pushd /tmp && git clone https://github.com/zigtools/zls
pushd zls && git checkout 0.14.0
zig build -Doptimize=ReleaseSafe 
sudo mv zig-out/bin/zls /usr/local/bin/ && popd
rm -rf /tmp/zls && popd
