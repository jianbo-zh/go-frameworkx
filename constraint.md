

命名约束

目录如果是一个包，则不能带 s 后缀

严格分层设计： api/svc/repo  task/svc/repo 

api > task > svc > repo

很关注的东西且各个层次都需要用到则，提升为公共的 比如：用户类型 或 日志定义


命名规则：
需要导出的包，目录需要带类型后缀，如： mod, svc, repo, mdl, cmn 标识类型，来区分角色
用户定义的类型，目录或文件名都为： utype
日志器，目录命名：logger


父目录的库可以导入子目录的库，而子目录不能直接使用父目录的库，所以公共库需要新建目录（否则引用循环了）

目录下的同名文件作为对应包的是入口文件

库名和目录同名，库名不能带s后缀

同类型的公共文件或目录使用common命名


----------------- 通用包命名（包名） -----------------
utype  用户定义类型
udefine  用户定义变量或常量
logger 日志器

----------------- 后缀约定（包名后缀） -----------------
mod  module
svc  service
svr  server
repo repository 
mdl  mdl
cmn  common

----------------- 变量名约定 -----------------
循环变量：
i       index
k       key
v       value

一个名字的声明和使用之间的距离越大，这个名字的长度就越长


----------------- 文件命名约定 -----------------
文件名称用下划线分割
最后一个单词表示该文件类型或用途


----------------- 编码规范 -----------------
每行代码最好不要超过100个字符，绝不超过120个字符
vscode设置： rulers: [100, 120]


init 为相同包，但多个文件共用的 变量或常量，或者被其他包引用的 常量或变量？
setup  指当前包使用前，需要明确初始化的方法，一般有 Setup方法
包名同名的文件，为该包的主要方法定义