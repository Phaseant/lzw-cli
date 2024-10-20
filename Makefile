.PHONY: all

TEXT_PATH=./text.txt
DICT_PATH=./o_dict.txt
OUTPUT_PATH=encoded.txt

dict:
	go run main.go generate-dict --text-path $(TEXT_PATH) --dictionary-path $(DICT_PATH)

encode:
	go run main.go encode --text-path $(TEXT_PATH) --dictionary-path $(DICT_PATH) --output-path $(OUTPUT_PATH)

decode:
	go run main.go decode --encoded-text-path $(OUTPUT_PATH) --dictionary-path $(DICT_PATH)