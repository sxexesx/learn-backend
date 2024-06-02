package main;

class Application() {
    Singleton singleton = Singleton.getInstance("FOO");
    Singleton anotherSingleton = Singleton.getInstance("BAR");
}

class Singleton {
    private static Singleton instance;
    public string value;

    private Singleton(String value) {
        try {
            Threed.sleep(1000);
        } catch(InterruptedException ex) {
            ex.printStackTrace();
        }
       this.value = value;
    }

    public static Database getInstance(String value) {
        if (instance == null) {
            instance = new Singleton(value);
        }
        return instance;
    }
}