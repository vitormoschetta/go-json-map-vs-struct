# JSON Converto to Map or Struct?

Em Go podemos escolher entre `map[string]interface{}` e objetos estruturados (structs) ao lidar com deserialização de JSON. Qual é a melhor escolha?

Aqui estão algumas considerações que podem ajudá-lo a tomar a decisão:

## 1. Mapas
- Usado quando você não conhece a estrutura dos dados que serão enviados na solicitação HTTP, como em APIs genéricas que recebem payloads variáveis.  
- Pode ser flexível, mas perde o tipo estático de Go, o que significa que você precisa fazer verificações de tipo em tempo de execução para acessar os dados.  
- Pode ser mais difícil de manter e depurar, especialmente em projetos maiores, devido à falta de estrutura definida.  

## 2. Structs
- Usado quando você conhece a estrutura dos dados que serão enviados na solicitação HTTP, como em APIs que recebem payloads específicos.
- Oferecem benefícios de tipo estático, permitindo que o compilador verifique os tipos em tempo de compilação.
- Tornam o código mais legível e ajudam na manutenção, pois a estrutura é explicitamente definida.