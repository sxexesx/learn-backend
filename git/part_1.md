# GIT

**Git** - система версионирования кода  


**Коммит (англ. commit)** – слепок всех отлеживаемых гитом (индексированных) файлов. Каждый коммит имеет родительский 
коммит – предыдущее состояние файла, а совокупность всех коммитов формирует дерево коммитов. Название каждого коммита 
уникально и представляет собой _хэш-сумму коммита_.

Если файл не изменился при переходе от одного коммита к другому, то гит укажет на ссылку файла в предыдущем состоянии. 
Т.е. физически файл храниться один раз, а гит хранит слепки/срезы.

![img.png](src/img.png)

### Стадии работы с git:

1. Редактирование файла – изменение файла в любом текстовом редакторе.
2. Добавление файла в staging area – область с изменениями, которые будут включены в следующий коммит.
3. _Сохранение коммита_ – перманентное сохранение всех собранных в staging area изменений.

![img.png](src/img02.png)

### Основные команды

Для инициализации git в директории, за содержимым которой нужно следить, используется команда `git init`. В результате 
работы команды создается скрытая папка `.git `, в которой и будут храниться данные о всех коммитах.  

`git add` - добавить изменения в файле в staging area  
`git status` - проверкa статуса файлов  
`git commit` - сделать коммит  
`git log` - показать все коммиты, с флагом `-graph` - дерево коммитов  
`git diff --staged` - выводит изменения, добавленные в staging area  