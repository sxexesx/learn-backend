public abstract class Shape() {
    public int x;
    public int y;
    public String color;

//     constructor

    public abstract Shape clone();
}

public class Circle extends Shape {
    public int radius;

//      constructor

    @Override
    public Shape clone() {
        return new Circle(this);
    }
}

// client
public class App {
    public static void main() {
        List<Shape> shapes = new ArrayList<>();
        List<Shape> shapesCopy = new ArrayList<>();

        Circle c = new Circle();
        c.x = 10;
        c.y = 20;
        c.radius = 15;
        circle.color = "red";

        Circle anotherC =  (Circle) c.clone()

    }
}