# Boost Retrieval car file

Two types of files are generated during the sealing process of Filecoin,  respectively Unsealed files and CAR original files. Now booster-http retrieval is mainly for unsealed files that are required to be converted into CAR files to complete the retrieval. We created another method to suport direct retrieval for Car original files, which reduces  the comsumption of system resources. Users can only keep car original files, this improvement saves  mass valuable storage resources.

## Functional transformation list
- booster-http set up car-urls parameters，car server urls, split using semicolon
- lotus storage tryReadUnsealedPiece add method to read car original file
- lotus-car-storage Newly add car original files remote reading server