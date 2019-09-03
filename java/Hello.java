import java.util.Scanner;
public class Hello {
    public static void main(String[] args) {
        String fruit = "apple";
        int opt = switch (fruit) {
            case "apple" -> 1;
            case "pear", "mango" -> 2;
            default -> 0;
        };
        System.out.println("opt = " + opt);
        // Scanner scanner = new Scanner(System.in);
        // System.out.println("input your name:");
        // String name = scanner.nextLine();
        // System.out.println("input your age:");
        // int age = scanner.nextInt();
        // System.out.printf("hi, %s you are %d\n ", name, age);
        // String s = "hello";
        // String t = s;
        // s = "world";
        // System.out.println(t);
    }
}
