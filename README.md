# guilib

A no_std monochrome raster graphics library


## WebAssembly Demo

Hosted: https://samblenny.github.io/guilib/examples/mkwasm/www/

Local (requires ruby and make):

```
cd examples/mkwasm
# Build wasm32-unknown-unknown binary and copy to examples/mkwasm/www
make install
# Start a dev webserver on http://localhost:8000
cd www
ruby webserver.rb
```


## Code Generation Notes

The rust source code for bitmap fonts in `src/fonts/` was generated by
javascript in the web page contained in `www_codegen`.

Procedure to update source code for the bitmap fonts:

1. Update bitmap glyphs in `www_codegen/img/*.png`

2. Update character map in `www_codegen/main.js`

3. Run `www_codegen/webserver.rb`

4. Load localhost:8000/ in a modern browser (ES6 support)

5. Select a png file from the "bitmap" pulldown list (the rust source code in
   the box at the bottom of the page should update)

6. Copy the rust source code from the box at the bottom of the page and paste
   it into the appropriate file in `src/fonts/*.rs`

7. If needed, repeat steps 5 and 6 for the other png files


## Credits

See [CREDITS.md](CREDITS.md)


## License

Dual licensed under the terms of [Apache 2.0](LICENSE-APACHE) or
[MIT](LICENSE-MIT), at your option.
