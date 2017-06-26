package se.dw.aoc2016.day24;

public class Position {
    int x, y, distance = 0;

    public Position(int x, int y) {
        this.x = x;
        this.y = y;
    }

    @Override
    public int hashCode() {
        return Integer.parseInt(x + "000" + y);
    }

    @Override
    public boolean equals(Object obj) {
        Position p2 = (Position) obj;
        if (this.x == p2.getX() && this.y == p2.getY()) {
            return true;
        }
        return false;
    }

    public Position(int x, int y, int distance) {
        this.x = x;
        this.y = y;
        this.distance = distance;
    }

    @Override
    public String toString() {
        return "Position(" + x + "," + y + ") d=" + distance;
    }

    public int getY() {
        return y;
    }

    public int getX() {
        return x;
    }

    public void setDistance(int distance) {
        this.distance = distance;
    }

    public int getDistance() {
        return distance;
    }

}