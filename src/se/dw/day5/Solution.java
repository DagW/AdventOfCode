package se.dw.day5;

import javax.xml.bind.DatatypeConverter;
import java.security.MessageDigest;
import java.security.NoSuchAlgorithmException;

/**
 * My solution to the day5 problem
 * http://adventofcode.com/2016/day/5
 */
public class Solution {

    public final static String input = "ffykfhsq";
    public String password = "";

    public Solution(String input) throws NoSuchAlgorithmException {

        MessageDigest md5 = MessageDigest.getInstance("MD5");

        /*
        Without a way to predict what hashes will start with 00000
        we have to brute force
         */

        for (int i = 0; i < Integer.MAX_VALUE; i++) {

            String toHash = input + i;

            byte[] hash = md5.digest(toHash.getBytes());
            String hex = DatatypeConverter.printHexBinary(hash);

            //When we find the right hash, add the 6th char to the password
            if (hex.startsWith("00000")) {
                char sixthChar = hex.charAt(5);
                password += sixthChar;
            }

            if (password.length() >= 8) {
                break;
            }

        }

        System.out.println("Solution: " + password);

    }

    public static void main(String[] args) {

        try {
            new Solution(input);
        } catch (NoSuchAlgorithmException e) {
            e.printStackTrace();
        }

    }

}