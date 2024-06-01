interface Builder {
    string setOptionA(string a);
    string setOptionB(string b);
    string setOptionC(string c);
}

class ConcreteBuilderA implements Builder {
    private string a;
    private string b;
    private string c;

    @Override
    string setOptionA(string a) {
        this.optionA = a
    }

    @Override
    string setOptionA(string b) {
        this.optionB = b
    }

    @Override
    string setOptionC(string c) {
        this.optionB = c
    }

    public ConcreteResultA getResult() {
        return new ConcreteResultA(a, b, c);
    }
}

class ConcreteBuilderB implements Builder {
    private string a;
    private string b;
    private string c;

    @Override
    string setOptionA(string a) {
        this.optionA = a
    }

    @Override
    string setOptionA(string b) {
        this.optionB = b
    }

    @Override
    string setOptionC(string c) {
        this.optionB = c
    }

    public ConcreteResultB getResult() {
        return new ConcreteResultB(a, b, c);
    }
}

public class ConcreteResultA {
    public string optionA;
    public string optionB;
    public string optionC;

    public ConcreteResultA(string a, string b, string c) {
        this.a = a;
        this.b = b;
        this.c = c;
    }
//     getters, setters
}

public class ConcreteResultB {
    public string optionA;
    public string optionB;
    public string optionC;

    public ConcreteResultB(string a, string b, string c) {
            this.a = a;
            this.b = b;
            this.c = c;
        }

}

// директор
public class Director {
    public void constructProduct(Builder builder) {
        builder.setOptionA("option-1");
        builder.setOptionB("option-2");
        builder.setOptionC("option-3");
    }
}

// client
public class App {
    public static void main(args...) {
        Director d = new Director();

        Builder builderA = new ConcreteBuilderA();
        d.constructProduct(builderA);

        ConcreteResultA result = builderA.getResult();

//      -------------------------
        Builder builderB = new ConcreteBuilderB();
        d.constructProduct(builderB);

        ConcreteResultB result1 = builderB.getResult();
    }
}