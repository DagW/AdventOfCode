package se.dw.day3;

public class Triangle {

    private final int side0, side1, side2;

    public Triangle(String triangle) {
        String[] stringSides = triangle.trim().split("\\s+");

        side0 = Integer.parseInt(stringSides[0]);
        side1 = Integer.parseInt(stringSides[1]);
        side2 = Integer.parseInt(stringSides[2]);
    }

    public boolean isValid() {

        //All two side sums must be larger than the third side
        if (side0 + side1 <= side2) {
            return false;
        } else if (side1 + side2 <= side0) {
            return false;
        } else if (side2 + side0 <= side1) {
            return false;
        }
        return true;

    }
}