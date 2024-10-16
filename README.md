# fixmultifiles
# fix multiple files - get rid of viruses
This software (written in GO) was designed to get rid of one javascript virus that placed itself at the end of every javascipt file in a WordPress installation
Over 3500 files were infected.  WordFence managed to elimiate closs to 3000, but 650 remained infected.  WordFence didn't have access to the correct 
version of the file.

So I wrote this software.  Very simple - you will see it uses 4 virus signatures to make sure I'm eliminating a virus.  You can make this work for multiple
viruses by placing signatures in a file, having the software read in all the signatures to start, and then going down the list of signatures.  For an example of the virus see virus.js and search for the string ";if(typeof zqxq===" which is the start of the virus.

Written by Maurice Bizzarri, Bizzarri Software (now retired but still coding).
October, 2024
Version 0.2 release October 16, 2024

