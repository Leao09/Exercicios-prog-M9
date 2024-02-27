## Ponderada 3- Relatório sobre simulações de ataque 

# Introdução

Este relatório tem como objetivo fornecer uma análise das possibilidades de ataques e invasões, considerando os princípios fundamentais de segurança, como confiabilidade, integridade e disponibilidade. A seguir, apresentamos descrições detalhadas das potenciais ameaças.
## Confiabilidade
### Ataques de Negação de Serviço (DoS):

- Atacantes podem realizar ataques de negação de serviço inundando o servidor MQTT com uma quantidade massiva de solicitações, tornando-o inacessível. Por exemplo, um ataque DoS pode envolver milhões de requisições simultâneas, esgotando os recursos do servidor e prejudicando a conectividade dos dispositivos.
### Saturação de Banda Larga:

- Atacantes podem empregar técnicas como flooding, enviando volumes excessivos de dados pela rede MQTT para consumir toda a largura de banda disponível. Isso pode resultar em congestionamento da rede, levando a atrasos nas mensagens e interrupções na comunicação entre os dispositivos.
### Exaustão de Recursos:

- Atacantes podem explorar vulnerabilidades no protocolo MQTT para iniciar ataques de exaustão de recursos, como o envio repetitivo de mensagens de grande volume. Esse tipo de ataque pode sobrecarregar a capacidade do servidor MQTT, resultando em tempos de resposta mais lentos e eventual falha do sistema.

## Integridade
### Modificação de Mensagens Durante a Transmissão:

- Atacantes podem usar técnicas de Man-in-the-Middle para interceptar e modificar mensagens MQTT durante a transmissão. Por exemplo, um invasor pode alterar comandos enviados de um dispositivo para outro, causando a execução de ações não autorizadas.

### Replay Attacks (Ataques de Repetição):

- Atacantes podem gravar e retransmitir mensagens anteriores, como um comando de desligar um dispositivo. Isso pode resultar na repetição da ação, comprometendo a integridade das operações e levando a resultados inesperados.

### Injeção de Mensagens Falsas:

- Atacantes podem injetar mensagens falsas na rede MQTT, como falsas atualizações de sensores. Isso pode levar a decisões incorretas, como acionar alarmes desnecessários ou tomar ações não apropriadas com base em dados adulterados.

## Disponibilidade

### Ataques de Negação de Serviço Distribuído (DDoS):

- Atacantes podem orquestrar ataques DDoS coordenando uma rede distribuída para sobrecarregar o servidor MQTT. Por exemplo, um botnet pode ser utilizado para gerar tráfego massivo, levando a uma indisponibilidade prolongada dos serviços.

### Exaustão de Conexões:

- Atacantes podem criar um grande número de conexões simultâneas, excedendo a capacidade do servidor MQTT. Isso pode resultar na rejeição de conexões legítimas, tornando o serviço inacessível para dispositivos autorizados.

### Ataques de Sobrecarga de Filas:

- Atacantes podem publicar mensagens em alta frequência, enchendo as filas de tópicos e dificultando a capacidade do sistema de processar novas mensagens. Isso pode resultar em atrasos significativos na entrega de mensagens críticas.

## Conclusão

Adotar a perspectiva de um invasor proporciona insights valiosos para abordar de maneira objetiva e assertiva as vulnerabilidades de um projeto em relação à segurança de dados. Compreender as possíveis situações de ataque e as ferramentas que um invasor pode utilizar é fundamental para a criação eficiente de barreiras de segurança. A implementação de fatores de redundância em funcionalidades críticas é essencial para mitigar os impactos de vazamentos de informações ou ataques bem-sucedidos. Essa abordagem proativa é vital para fortalecer a segurança de uma aplicação em um cenário cada vez mais complexo e dinâmico.