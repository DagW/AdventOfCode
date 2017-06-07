package se.dw.day7;

import java.util.ArrayList;

/**
 * Created by dag on 2017-06-07.
 */
public class IPv7Address {

    //ABBA is 2 chars and the 2chars repeated backwards, cannot be 4 of the same chars
    //Rules to support TLS
    //#1 No ABBA's in within square brackets
    //#2 Atleast one ABBA in other parts

    private ArrayList<String> addressParts = new ArrayList<>();
    private ArrayList<String> hypernetSequences = new ArrayList<>();

    public IPv7Address(String address) {

        String[] parts = address.split("\\W+");

        for (int i = 0; i < parts.length; i++) {
            if (i % 2 == 0) {
                addressParts.add(parts[i]);
            } else {
                hypernetSequences.add(parts[i]);
            }
        }

    }

    public boolean containsAbba(String input) {

        char[] chars = input.toCharArray();

        for (int i = 0; i < chars.length; i++) {

            char[] abba = new char[4];

            //Populate the abba array with the next for chars, if they exist
            for (int j = 0; j < 4 && j < chars.length - i; j++) {
                abba[j] = chars[i + j];
            }

            //If any of the char items are un-initialized - there cant be an abba
            for (char c : abba) {
                if (c == '\u0000') {
                    return false;
                }
            }

            //Check the abba with the two rules
            //ABBA is 2 chars and the 2chars repeated backwards, cannot be 4 of the same chars
            if ((abba[0] == abba[3] && abba[1] == abba[2]) &&
                    (abba[0] != abba[1])) {
                return true;
            }

        }

        return false;
    }

    public boolean supportsTLS() {

        //If the bracketed text contains an abba, fail
        for (String str : hypernetSequences) {
            if (containsAbba(str)) {
                return false;
            }
        }

        //If the other texts contains one, success
        for (String str : addressParts) {
            if (containsAbba(str)) {
                return true;
            }
        }

        return false;
    }

}
