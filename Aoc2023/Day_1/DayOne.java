import java.io.File;
import java.io.FileNotFoundException;
import java.util.Scanner;

class Numbers {
    int first;
    int firstIndex;
    int lastIndex;
    int last;

    int firstWord;
    int firstWordIndex;
    int lastWord;
    int lastWordIndex;

    public Numbers() {
        first = -1;
        last = -1;
        firstWord = -1;
        lastWord = -1;
        firstIndex = 1000000;
        lastIndex = -1;
        firstWordIndex = -1;
        lastWordIndex = -1;
    }

    public void setFirst(int value, int index) {
        first = value;
        firstIndex = index;
    }

    public void setLast(int value, int index) {
        last = value;
        lastIndex = index;
    }

    public void setFirstWord(int value, int index) {
        firstWord = value;
        firstWordIndex = index;
    }

    public void setLastWord(int value, int index) {
        lastWord = value;
        lastWordIndex = index;
    }

    public int getFirstWord() {
        return firstWord;
    }

    public int getLastWord() {
        return lastWord;
    }

    public int getFirst() {
        return first;
    }

    public int getLast() {
        return last;
    }

    public int getFirstIndex() {
        return firstIndex;
    }

    public int getLastIndex() {
        return lastIndex;
    }

    public int getFirstWordIndex() {
        return firstWordIndex;
    }

    public int getLastWordIndex() {
        return lastWordIndex;
    }

    public int result() {
        return 10 * first + last;
    }

    public void compute() {

        System.out.println("-........");
        System.out.println("Fir: " + getFirst() + " ix: " + getFirstIndex());
        System.out.println("Last: " + getLast() + " ix: " + getLastIndex());
        System.out.println("FirWord: " + getFirstWord() + " ix: " +
                getFirstWordIndex());
        System.out.println("LastWord: " + getLastWord() + " ix: " +
                getLastWordIndex());

        // preveri za spodnjo stevilko

        if (getFirst() == -1) {
            setFirst(getFirstWord(), getFirstIndex());
        } else if (getFirstWord() != -1 && getFirstWordIndex() < getFirstIndex()) {
            setFirst(getFirstWord(), getFirstWordIndex());
        }

        // preveri za zgornjo stevilko
        if (getLastWord() == -1 && getLast() == -1) {
            setLast(getFirst(), getFirstIndex());
        } else if (getLastWord() != -1 && getLastIndex() < getLastWordIndex()) {
            setLast(getLastWord(), getLastWordIndex());
        }
    }
}

class DayOne {

    public static int wordCheck(int i, String line, boolean max) {
        switch (i) {
            case 1:
                if (max)
                    return line.indexOf("one");
                return line.lastIndexOf("one");
            case 2:
                if (max)
                    return line.indexOf("two");
                return line.lastIndexOf("two");
            case 3:
                if (max)
                    return line.indexOf("three");
                return line.lastIndexOf("three");

            case 4:
                if (max)
                    return line.indexOf("four");
                return line.lastIndexOf("four");

            case 5:
                if (max)
                    return line.indexOf("five");
                return line.lastIndexOf("five");

            case 6:
                if (max)
                    return line.indexOf("six");
                return line.lastIndexOf("six");

            case 7:
                if (max)
                    return line.indexOf("seven");
                return line.lastIndexOf("seven");

            case 8:
                if (max)
                    return line.indexOf("eight");
                return line.lastIndexOf("eight");

            case 9:
                if (max)
                    return line.indexOf("nine");
                return line.lastIndexOf("nine");

            default:
                return -1;
        }
    }

    public static Numbers eval(String line) {
        Numbers res = new Numbers();

        for (int i = 0; i < line.length(); i++) {
            if (Character.isDigit(line.charAt(i))) {
                if (res.getFirstIndex() > i) {
                    res.setFirst(Character.getNumericValue(line.charAt(i)), i);
                }

                if (res.getLastIndex() < i) {
                    res.setLast(Character.getNumericValue(line.charAt(i)), i);
                }
            }
        }

        for (int i = 1; i < 10; i++) {
            int wc = wordCheck(i, line, false);
            int wc2 = wordCheck(i, line, true);
            // System.out.println("St: " + i + " wc: " + wc);

            if (wc2 != -1 && wc2 < res.getFirstIndex()) {
                res.setFirst(i, wc2);
            }

            if (wc != -1 && wc > res.getLastIndex()) {
                res.setLast(i, wc);
            }

        }

        // nastavi najmanjse in najvecje
        // res.compute();

        return res;

    }

    public static void main(String[] args) throws FileNotFoundException {
        File inp = new File("Input.txt");

        Scanner sc = new Scanner(inp);

        int sum = 0;

        while (sc.hasNextLine()) {
            String data = sc.nextLine();
            Numbers res = eval(data);
            sum += res.result();
            // System.out.printf("%s -> Fir: %d, Last: %d, RES: %d \n", data,
            // res.getFirst(),
            // res.getLast(), res.result());
        }

        System.out.println("Result: " + sum);

        sc.close();
    }
}