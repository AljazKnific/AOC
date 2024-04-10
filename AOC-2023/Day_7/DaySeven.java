import java.io.File;
import java.util.Arrays;
import java.util.HashMap;
import java.util.LinkedList;
import java.util.Scanner;

class Hand {
    String cards;
    int val;
    int type;
    int max;

    // Which placement are we in

    public Hand(String cards, int val) {
        this.cards = cards;
        this.val = val;
        this.type = detType(cards);
    }

    public int detType(String cards) {
        HashMap<Character, Integer> map = new HashMap<Character, Integer>();
        max = 0;
        int joker = 0;

        for (int i = 0; i < 5; i++) {
            char c = cards.charAt(i);

            if (c == 'J') {
                joker++;
                continue;
            }

            Integer x = map.get(c);
            if (x == null) {
                map.put(cards.charAt(i), 1);
                if (1 > max) {
                    max++;
                }
            } else {
                map.replace(cards.charAt(i), x, x + 1);
                if (x + 1 > max) {
                    max++;
                }
            }
        }

        if (joker == 0) {

            switch (map.size()) {
                case 5:
                    return 0;
                case 4:
                    return 1;
                case 3:
                    if (max == 2) {
                        return 2;
                    }
                    return 3;
                case 2:
                    if (max == 3)
                        return 4;
                    return 5;
                case 1:
                    return 6;
                default:
                    return -1;
            }
        } else {
            switch (max + joker) {
                case 2:
                    return 1;
                case 3:
                    if (map.size() == 2) {
                        return 4;
                    }
                    return 3;
                case 4:
                    return 5;
                case 5:
                    return 6;
                default:
                    return -1;
            }
        }

    }

    public void izpis() {
        System.out.println(cards + " -> val: " + val + "  type: " + type + " max: " + max);
    }

}

class DaySeven {

    public static int getValueOfCard(char c) {
        if (Character.isDigit(c))
            return Character.getNumericValue(c);

        int x = 0;
        switch (c) {
            case 'A':
                x = 14;
                break;
            case 'K':
                x = 13;
                break;
            case 'Q':
                x = 12;
                break;
            case 'J':
                // partTwo- joker
                x = 1;
                break;
            case 'T':
                x = 10;
                break;
            default:
                x = -1;
                break;
        }

        return x;
    }

    public static boolean isBetter(Hand h1, Hand h2) {
        // System.out.println(h1.cards + " : " + h2.cards);
        for (int i = 0; i < 5; i++) {
            int a = getValueOfCard(h1.cards.charAt(i));
            int b = getValueOfCard(h2.cards.charAt(i));

            if (a < b) {
                // System.out.println("A: " + a + " B: " + b);
                return false;
            } else if (a > b) {
                return true;
            }

        }

        return true;
    }

    public static void handsRightPos(LinkedList<Hand> list, Hand h) {
        int i = 0;

        while (i < list.size()) {
            Hand temp = list.get(i);

            if (isBetter(temp, h)) {
                break;
            }
            i++;
        }

        list.add(i, h);
    }

    public static void main(String[] args) {

        try {

            LinkedList<LinkedList<Hand>> list = new LinkedList<>();
            for (int i = 0; i < 7; i++) {
                list.add(new LinkedList<>());

            }

            Scanner sc = new Scanner(new File("Input.txt"));

            while (sc.hasNextLine()) {
                String[] data = sc.nextLine().split(" ");
                Hand h = new Hand(data[0], Integer.parseInt(data[1]));
                handsRightPos(list.get(h.type), h);
            }

            int st = 1;
            int sum = 0;
            for (int i = 0; i < list.size(); i++) {
                for (int j = 0; j < list.get(i).size(); j++) {
                    Hand h = list.get(i).get(j);

                    // h.izpis();

                    sum += (st * h.val);
                    st++;
                }
            }

            System.out.println("Sum: " + sum);

        } catch (Exception e) {
            System.err.println(e);
        }
    }
}