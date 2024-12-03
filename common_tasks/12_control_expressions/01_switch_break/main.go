loop:
for {
        switch expr {
        case foo:
                if condA {
                        doA()
                        break // like 'goto A' // те брейк без метки - это выход из свича (не из цикла!!!)
                }

                if condB {
                        doB()
                        break loop // like 'goto B'   // брейк с меткой - выход из цикла и дальшейшее выполнение кода после цикла                      
                }

                doC()
        case bar:
                // ...
        }
A:
        doX()
        // ...
}

B:
doY()
// ....
