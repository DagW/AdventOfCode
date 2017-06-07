package se.dw.day9;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * My solution to the day9 problem
 * http://adventofcode.com/2016/day/9
 */
public class Solution {

    public Solution(String input) {

        Decompressor decompressor = new Decompressor(input);
        System.out.println("Before decompress: " + input.length());

        decompressor.decompressData();

        System.out.println("Solution: " + decompressor.decompressedData.length());

    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day9.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);
    }

}