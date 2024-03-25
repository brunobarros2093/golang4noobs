# Herança e Composição 

Em Go, não há suporte direto para herança de classes como em linguagens orientadas a objetos tradicionais. No entanto, é possível alcançar comportamentos semelhantes usando composição e interfaces.

## Composição

A composição é um princípio fundamental em Go, onde uma struct pode incluir outras structs como campos.

```
type Animal struct {
    nome string
    idade int
}

type Cachorro struct {
    animal Animal
    raca   string
}

```

Aqui, Cachorro tem um campo animal do tipo Animal, permitindo que ele acesse os campos e métodos de Animal. No entanto, isso não é herança no sentido clássico, mas sim uma forma de reutilização de código.

## Interfaces

Em Go, a herança de tipo é geralmente feita através de interfaces. Se uma struct implementa todos os métodos de uma interface, ela implicitamente implementa essa interface. 

Você pode ver mais sobre [Interfaces](https://github.com/brunobarros2093/golang4noobs/tree/main/15%20-%20interfaces) aqui. 

## Incorporação de tipos

Go também oferece a capacidade de incorporar tipos diretamente em outras structs, o que é semelhante à herança de campos em outras linguagens. No entanto, isso também é uma forma de composição e não herança no sentido tradicional de orientação a objetos.

```
type Animal struct {
    nome string
    idade int
}

type Cachorro struct {
    Animal
    raca string
}
```
Aqui, Cachorro incorpora o tipo Animal, o que significa que todos os campos e métodos de Animal estão disponíveis diretamente em Cachorro.

Embora Go não suporte herança de classes como em linguagens orientadas a objetos tradicionais, esses conceitos permitem alcançar muitos dos mesmos objetivos, promovendo a reutilização de código e uma estrutura de código limpa e eficiente.