package se.dw.day7;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;

/**
 * My solution to the day7 problem
 * http://adventofcode.com/2016/day/7
 */
public class Solution {

    public Solution(String input) {

        int supportingTlsCount = 0;

        for (String address : input.split("\n")) {

            IPv7Address iPv7Address = new IPv7Address(address);
            if(iPv7Address.supportsTLS()){
                supportingTlsCount ++;
            }

        }

        System.out.println("Solution: "+supportingTlsCount);

    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/day7.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}