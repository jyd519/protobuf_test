SRC_DIR=.
DST_DIR=.

gen:
		mkdir -p tutorial
		protoc -I=${SRC_DIR} --python_out=${DST_DIR} ${SRC_DIR}/addressbook.proto
		protoc -I=${SRC_DIR} --go_out=tutorial ${SRC_DIR}/addressbook.proto


# vim: set noet ts=2 sw=2
