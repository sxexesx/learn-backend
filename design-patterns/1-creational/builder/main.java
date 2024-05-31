interface Builder {
    string setOptionA(string a);
    string setOptionB(string b);
    string setOptionC(string c);
}

class ConcreteBuilderA implements Builder {
    private ConcreteResultA resultA;

    @Override
    string setOptionA(string a) {
        resultA.optionA = a
    }

    @Override
    string setOptionA(string b) {
        resultA.optionB = b
    }

    @Override
    string setOptionC(string c) {
        resultA.optionB = c
    }

    public ConcreteResultA getResult() {

    }
}

class ConcreteBuilderB implements Builder {
}

interface Result {
    public string optionA;
    public string optionB;
    public string optionC;
}

class ConcreteResultA implements Result {
    public string optionA;
    public string optionB;
    public string optionC;
}

class ConcreteResultB implements Result {
    public string optionA;
    public string optionB;
    public string optionC;
}


