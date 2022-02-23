
This is a little experiment to see what happens when I change a `oneof` to a regular field.

* `generate.sh` runs the proto commands to generate the code for the contracts
* `main1` creates a message in first contract and writes it into `msg.bin` file
* `main2` reads the message from the file

I tried using both contracts in a single main but Proto complained about that.
