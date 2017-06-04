package se.dw.day4;

import java.util.Arrays;
import java.util.Comparator;
import java.util.HashMap;

public class Room {

    private HashMap<Character, Integer> letters = new HashMap<>();
    private String checksum;
    private int roomId;


    public Room(String roomString) {

        //The first part containing the letters to be indexed
        String firstPart = roomString.substring(0, roomString.lastIndexOf("-"));
        firstPart = firstPart.replace("-", "");

        //Extract letter counts
        char[] letterArr = firstPart.toCharArray();
        for (int i = 0; i < letterArr.length; i++) {
            char letter = letterArr[i];
            if (letters.containsKey(letter)) {
                letters.put(letter, letters.get(letter) + 1);
            } else {
                letters.put(letter, 1);
            }
        }

        //This is the roomId[checksum]-part
        String lastPart = roomString.substring(roomString.lastIndexOf("-"));
        checksum = lastPart.substring(lastPart.indexOf("[") + 1, lastPart.lastIndexOf("]"));
        roomId = Integer.parseInt(lastPart.substring(0, lastPart.indexOf("[")));

    }

    public String calculateChecksum() {
        String checksum = "";

        Object[] keys = letters.keySet().toArray();
        Arrays.sort(keys, new Comparator<Object>() {

            @Override
            public int compare(Object o1, Object o2) {
                Character c1 = (Character) o1;
                Character c2 = (Character) o2;

                if (letters.get(c1) == letters.get(c2))
                    return c1.compareTo(c2);
                else
                    return letters.get(c2) - letters.get(c1);
            }
        });
        for (int i = 0; i < keys.length && i < 5; i++) {
            checksum += keys[i];
        }

        return checksum;
    }

    public String getChecksum() {
        return checksum;
    }

    public int getRoomId() {
        return roomId;
    }

    public boolean isChecksumValid() {
        return checksum.equals(calculateChecksum());
    }
}