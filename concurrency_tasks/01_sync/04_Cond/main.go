









// sync.Cond — это тип, представляющий условную переменную и используемый для координации горутин, позволяя им дождаться, пока определенное условие станет истинным, прежде чем продолжить.

// Тип sync.Cond позволяет создавать переменные условия и управлять ими. Имеет три основных метода:

// Wait() : этот метод заставляет вызывающую горутину ждать, пока другая горутина не сообщит об условной переменной. 
// Когда горутина вызывает Wait() , она снимает связанную блокировку и приостанавливает выполнение до тех пор, пока другая горутина не вызовет Signal() или Broadcast() для той же переменной sync.Cond .

// Signal() : этот метод пробуждает одну горутину, ожидающую условную переменную. 
// Если ожидают несколько горутин, пробуждается только одна из них. Выбор того, какая горутина будет пробуждена, произволен и не гарантирован.

// Broadcast() : этот метод пробуждает все горутины, ожидающие условной переменной. Когда вызывается Broadcast() , все ожидающие горутины пробуждаются и могут продолжить работу.

// Обратите внимание, что sync.Cond требует связанного sync.Mutex для синхронизации доступа к условной переменной.
