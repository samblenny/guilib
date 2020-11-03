#!/bin/bash
# This script downloads emoji graphics from twitter/twemoji and modifies them
# from color PNG format to monochrome PNG with dithered grayscale shading.
#
# To run the script, you need git, imagemagick, curl, pngcrush, and a bash
# compatible shell. The convert, montage, and mogrify commands are provided by
# imagemagick.
#
# The script will download 64 raw files out of the twitter/twemoji repo on
# GitHub. The full repo has directories with thousands of small files which can
# load slowly. Selective downloading is faster with less UI glitching.
#
# Twemoji License Notice:
# > Copyright 2019 Twitter, Inc and other contributors
# > Code licensed under the MIT License: http://opensource.org/licenses/MIT
# > Graphics licensed under CC-BY 4.0: https://creativecommons.org/licenses/by/4.0/
#
# Twemoji Source Code Link:
# - https://github.com/twitter/twemoji
#

# Set emoji codepoints used by Matrix protocol clients for SAS verification as
# this script's positional arguments
set -- 1f436 1f431 1f981 1f40e 1f984 1f437 1f418 1f430 1f43c 1f413 1f427 \
1f422 1f41f 1f419 1f98b 1f337 1f333 1f335 1f344 1f30f 1f319 2601 1f525 1f34c \
1f34e 1f353 1f33d 1f355 1f382 2764 1f600 1f916 1f3a9 1f453 1f527 1f385 1f44d \
2602 231b 23f0 1f381 1f4a1 1f4d5 270f 1f4ce 2702 1f512 1f511 1f528 260e 1f3c1 \
1f682 1f6b2 2708 1f680 1f3c6 26bd 1f3b8 1f3ba 1f514 2693 1f3a7 1f4c1 1f4cc

mkdir -p twemoji/color
mkdir -p twemoji/mono
cd twemoji/color
echo "Downloading color emoji PNG files..."
for cp in "$@"; do
    curl -L -O -C - https://raw.githubusercontent.com/twitter/twemoji/master/assets/72x72/$cp.png
done
echo "Converting color emoji to monochrome..."
for cp in "$@"; do
    convert -colorspace gray -auto-level +level 15%x75% -background white \
    -alpha remove -ordered-dither o4x4,2 $cp.png ../mono/$cp.png
done
echo "Making spritesheet montage..."
cd ../mono
montage *.png -background black -geometry 72x72+1+1 -tile 8x ../sas_emoji.png
cd ..
mogrify -bordercolor black -border 1 sas_emoji.png
pngcrush -ow -rem alla sas_emoji.png
echo "your monochrome emoji spritesheet file is at ./twemoji/sas_emoji.png"
