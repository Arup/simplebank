printf '%s\n' '.PHONY: createdb dropdb' \
'createdb:' \
'\tdocker exec -it postgres13.22 createdb --username=root --owner=root simple_bank' \
'' \
'dropdb:' \
'\tdocker exec -it postgres13.22 dropdb simple_bank' > Makefile
