package se.dw.day9;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Created by dag on 2017-06-07.
 */
public class Decompressor {

    String decompressedData = "";
    String compressedData;

    public Decompressor(String input) {
        compressedData = input;
    }

    public boolean decompressNext() {

        Pattern pattern = Pattern.compile("\\((\\d*[x]\\d*)\\)", Pattern.DOTALL);
        Matcher matcher = pattern.matcher(compressedData);

        if (matcher.find()) {
            String match = matcher.group(1);

            int numChars = Integer.parseInt(match.split("x")[0]);
            int repeats = Integer.parseInt(match.split("x")[1]);

            int start = compressedData.indexOf(match);
            int stop = compressedData.indexOf(match) + match.length();

            //Text before the match
            String pre = compressedData.substring(0, start - 1);
            //Text after the match
            String post = compressedData.substring(stop + 1, compressedData.length());
            //Text remaining to be decompressed after expansion
            String remaining = compressedData.substring(stop + 1 + numChars, compressedData.length());

            //Repeat the sequence
            String toExpand = post.substring(0, numChars);
            String expanded = "";
            for (int i = 0; i < repeats; i++) {
                expanded += toExpand;
            }

            decompressedData = decompressedData + pre + expanded;
            compressedData = remaining; //Remaining data

            return true;
        }

        return false;

    }

    public String decompressData() {

        //Decompress until we have no more matches
        while (decompressNext()) ;
        //Then the rest of the data is already expanded
        decompressedData += compressedData;

        return decompressedData;

    }
}
