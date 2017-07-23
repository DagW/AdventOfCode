package se.dw.aoc2016.day6;

import java.io.IOException;
import java.nio.file.Files;
import java.nio.file.Path;
import java.nio.file.Paths;
import java.util.HashMap;

/**
 * My solution to the day6 problem
 * http://adventofcode.com/2016/day/6
 * <p>
 * Count occurances of characters per column
 * The highest occurance is the sought character for that column
 */
public class Solution {

    private String solution = "";

    public Solution(String input) {

        String[] rows = input.split("\n");
        int numColumns = rows[0].length();

        String[] columns = new String[numColumns];
        //Initialize the items, otherwise the strings will start will "null" when we append chars
        for (int i = 0; i < columns.length; i++) {
            columns[i] = "";
        }

        //First we flip the rows to columns
        for (String row : rows) {
            char[] characters = row.toCharArray();
            for (int i = 0; i < characters.length; i++) {
                columns[i] += characters[i];
            }
        }

        for (String column : columns) {

            //Then we count occurrences
            HashMap<Character, Integer> occuranceMap = new HashMap<>();
            for (Character character : column.toCharArray()) {
                if (occuranceMap.containsKey(character)) {
                    occuranceMap.put(character, occuranceMap.get(character) + 1);
                } else {
                    occuranceMap.put(character, 1);
                }
            }

            //And add the highest occurring char to solution
            int columnMax = 0;
            Character columnMaxCharacter = null;
            for (Character character : occuranceMap.keySet()) {
                if (occuranceMap.get(character) > columnMax) {
                    columnMaxCharacter = character;
                    columnMax = occuranceMap.get(character);
                }
            }
            solution += columnMaxCharacter;

        }

        System.out.println("Solution: " + solution);

    }

    public String getSolution() {
        return solution;
    }

    public static void main(String[] args) throws IOException {

        Path path = Paths.get("res/2016/day6.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}