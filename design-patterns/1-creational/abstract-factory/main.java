
public interface Button {
    void paint();
}

public class MacButton implements Button {
    
    @Override
    void paint() {
        System.out.println("this is MacOS button")
    }
}

public class WinButton implements Button {

    @Override
    void paint() {
        System.out.println("this is Windows button") 
    }
}

public interface Checkbox {
    void paint();
}

public class MacCheckbox implements Checkbox {

    @Override
    void paint() {
        System.out.println("this is MacOS checkbox") 
    }
}

public class WinCheckbox implements Checkbox {

    @Override
    void paint() {
        System.out.println("this is Windows checkbox") 
    }
}

// factory
public interface AbstractFactory {
    Button createButton();
    Checkbox createCheckbox();
}

public abstract class MacFactory implements AbstractFactory {
    public Button createButton() {
        return new MacButton();
    }

    public Checkbox createCheckbox() {
        return new MacCheckbox();
    }
}

public abstract class WinFactory implements AbstractFactory {
    public Button createButton() {
        return new WinButton();
    }

    public Checkbox createCheckbox() {
        return new WinCheckbox();
    }
}

// client
class Application {
    private Button button;
    private Checkbox checkbox;

    public Application(AbstractFactory factory) {
        button = factory.createButton();
        checkbox = factory.createCheckbox
    }
}

public static void main() {
    public static void main() {
        config = SomeConfigurations();

        systemOS = System.getProperty("os.name").toLowerCase();

        AbstractFactory factory;
        if systemOS.equals("windows") {
            factory = new WinFactory();
        }
        if systemOS.equals("mac") {
            factory = new MacFactory();
        }

        app = new Application(factory);
        app.paint();
    }
}
