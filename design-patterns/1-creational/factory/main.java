public interface Button {
    void render();
    void onClick();
}

public class WindowsButton implements Button {
    public void render() {
        System.out.println("rendered Windows button");
    }

    public void onClick() {
        System.out.println("clicked Windows button");
    }
}

public class MacButton implements Button {
    public void render() {
        System.out.println("rendered MacOS button");
    }

    public void onClick() {
        System.out.println("clicked MacOS button");
    }
}

// factory
public abstract class Dialog {
    public void renderWindow() {
        Button btn = createButton();
        btn.render();
    }

    public abstract Button createButton();
}

public class WindowsDialog extends Dialog {
    
    @Override
    public Button createButton() {
        return new WindowsButton();
    }
}

public class MacDialog extends Dialog {
    
    @Override
    public Button createButton() {
        return new MacButton();
    }
}

// client
public class Demo {
    public static Dialog dialog;
    
    public static void main() {
        configure();
        runBusinessLogic();
    }

    static void configure() {
        if (System.getProperty("os.name").equals("Windows 10")) {
            dialog = new WindowsDialog();
        } else {
            dialog = mew MacDialog();
        }
    }

    static void runBusinessLogic() {
        dialog.renderWindow();
    }


}
