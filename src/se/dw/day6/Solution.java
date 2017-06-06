package se.dw.day6;

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

        HashMap<Character, Integer>[] occuranceMap = new HashMap[numColumns];

        for (String row : rows) {
            char[] characters = row.toCharArray();
            for (int i = 0; i < characters.length; i++) {
                Character character = row.toCharArray()[i];

                HashMap<Character, Integer> map = occuranceMap[i];
                if (map == null)
                    map = new HashMap<>();


                if (map.containsKey(character)) {
                    map.put(character, map.get(character) + 1);
                } else {
                    map.put(character, 1);
                }

                occuranceMap[i] = map;

            }
        }

        for (HashMap<Character, Integer> map : occuranceMap) {
            int columnMax = 0;
            Character columnMaxCharacter = null;
            for (Character character : map.keySet()) {
                if (map.get(character) > columnMax) {
                    columnMaxCharacter = character;
                    columnMax = map.get(character);
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

        Path path = Paths.get("res/day6.txt");
        String input = new String(Files.readAllBytes(path));
        new Solution(input);

    }

}