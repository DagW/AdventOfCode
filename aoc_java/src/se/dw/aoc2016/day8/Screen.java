package se.dw.aoc2016.day8;

/**
 * Created by dag on 2017-06-07.
 */
public class Screen {

    private boolean[][] screenPixels;

    public Screen(int width, int height) {

        screenPixels = new boolean[width][height];

        System.out.println("Created new screen " + screenPixels.length + "x" + screenPixels[0].length);

        //Init the screen with the pixels off
        for (int x = 0; x < width; x++) {
            for (int y = 0; y < height; y++) {
                screenPixels[x][y] = false;
            }
        }

    }

    public void parseCommand(String command) {

        String[] cmdParts = command.split("\\W+");

        switch (cmdParts[0]) {

            case "rect":

                int width = Integer.parseInt(cmdParts[1].split("x")[0]);
                int height = Integer.parseInt(cmdParts[1].split("x")[1]);

                turnOnRect(width, height);

                break;

            case "rotate":

                int num = Integer.parseInt(cmdParts[3]);
                int byPixels = Integer.parseInt(cmdParts[5]);

                switch (cmdParts[1]) {
                    case "column":
                        rotateColumn(num, byPixels);
                        break;
                    case "row":
                        rotateRow(num, byPixels);
                        break;
                }

                break;

        }

    }

    /*
     * Turns of all pixels in a rect
     * starting at the top left extending width,height
     */
    private void turnOnRect(int width, int height) {

        System.out.println("Turning on pixels " + width + "x" + height);
        for (int x = 0; x < width && x < screenPixels.length; x++) {
            for (int y = 0; y < height && y < screenPixels[0].length; y++) {
                screenPixels[x][y] = true;
            }
        }

    }

    private void rotateColumn(int column, int byPixels) {

        System.out.println("Rotating column " + column + " by " + byPixels);

        for (int i = 0; i < byPixels; i++) {

            boolean temp = screenPixels[column][screenPixels[column].length - 1];
            for (int y = screenPixels[column].length - 2; y >= 0; y--) {
                screenPixels[column][y + 1] = screenPixels[column][y];
            }
            screenPixels[column][0] = temp;

        }

    }

    private void rotateRow(int row, int byPixels) {

        System.out.println("Rotating row " + row + " by " + byPixels);

        for (int i = 0; i < byPixels; i++) {

            boolean temp = screenPixels[screenPixels.length - 1][row];
            for (int x = screenPixels.length - 2; x >= 0; x--) {
                screenPixels[x + 1][row] = screenPixels[x][row];
            }
            screenPixels[0][row] = temp;

        }

    }

    public void printScreen() {

        for (int y = 0; y < screenPixels[0].length; y++) {

            String row = "";
            for (int x = 0; x < screenPixels.length; x++) {
                if (screenPixels[x][y])
                    row += "# ";
                else
                    row += ". ";
            }

            System.out.println(row);

        }

    }

    public int getLitPixels() {

        int count = 0;

        for (int y = 0; y < screenPixels[0].length; y++) {
            for (int x = 0; x < screenPixels.length; x++) {
                if (screenPixels[x][y])
                    count++;
            }
        }

        return count;

    }

}
