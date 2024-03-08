<img alt="LSR Cloud" width="40%" src="https://lucianoromao.com.br/lsr.png">

# Estudo sobre Kafka

## Visão Geral 🔎
O Apache Kafka é uma plataforma de streaming distribuída que fornece uma maneira eficiente, tolerante a falhas e escalável para enviar, armazenar e processar fluxos de eventos em tempo real. Foi originalmente desenvolvido pela LinkedIn e posteriormente aberto para a comunidade de código aberto da Apache.

O Kafka é amplamente utilizado em arquiteturas de microsserviços, processamento de eventos em tempo real, integração de dados e outras aplicações que exigem um sistema de mensagens distribuído e altamente escalável.

## Principais conceitos do Apache Kafka
 * Tópicos (Topics): Os dados são publicados em tópicos, que são categorias ou feeds de mensagens.
 * Produtores (Producers): São responsáveis por publicar mensagens em tópicos.
 * Consumidores (Consumers): Consomem mensagens de tópicos. Os consumidores podem ser agrupados para dimensionamento e tolerância a falhas.
 * Broker Kafka: Cada nó no cluster Kafka é chamado de broker. Ele é responsável por armazenar dados e atender solicitações de produtores e consumidores.
 * Partições: Cada tópico é dividido em partições, permitindo a distribuição de dados e paralelismo no processamento.
 * ZooKeeper: O Apache Kafka depende do Apache ZooKeeper para coordenação e gerenciamento de configuração no cluster.
 
## Kafka e resiliencia 
O Apache Kafka é considerado resiliente por várias razões, principalmente devido à sua arquitetura distribuída e características de tolerância a falhas. Algumas das razões para a resiliência do Kafka incluem:

 * Replicação de Dados: O Kafka permite a replicação de dados entre os nós do cluster. Cada partição de tópico pode ter cópias em vários brokers. Isso oferece tolerância a falhas, pois, mesmo que um broker falhe, as réplicas das partições podem ser usadas para garantir a disponibilidade contínua dos dados.

 * Particionamento e Escalabilidade: As partições permitem a distribuição de carga e dados entre os brokers. Isso facilita a escalabilidade horizontal, permitindo adicionar mais brokers conforme necessário. Se um broker falha, as partições ainda podem ser manipuladas por outros brokers no cluster.

 * Persistência: Kafka armazena mensagens em disco, garantindo durabilidade e recuperação em caso de falhas. As mensagens são retidas por um período configurável, e os consumidores podem recuperar eventos passados mesmo após uma falha temporária.


## Garantia de entrega das mensagens
O Kafka também possuí configurações interessante de ACK que permitem que a entrega de mensagens seja resiliente, embora possamos ter que sacrificar um pouco da perfomance para isso.

Tipos de ACK:
 * 0 (zero): O produtor não espera nenhuma confirmação de recebimento. A mensagem é considerada entregue assim que é enviada para o tópico no broker. Isso proporciona uma baixa garantia de entrega, pois não há confirmação de que a mensagem foi recebida pelo broker.

 * 1 (um): O produtor aguarda a confirmação de recebimento do líder da partição. A mensagem é considerada entregue assim que o líder recebe a mensagem. Isso oferece um nível intermediário de garantia, pois o líder pode replicar a mensagem para os seguidores após o recebimento.

 * all (-1 ou "all"): O produtor espera a confirmação de recebimento de todos os brokers da partição antes de considerar a mensagem entregue. Isso oferece a maior garantia de entrega, pois todas as réplicas da partição devem receber a mensagem.

## Indepotência
O Kafka também oferece opções de indepotência, garantindo que mesmo em situações de falhas ou retransmissões, não ocorram efeitos colaterais indesejados, como a criação de mensagens duplicadas. 

 * A indepotência do produtor refere-se à capacidade de enviar mensagens de forma que, mesmo que ocorram falhas ou retransmissões, não haja efeitos colaterais indesejados, como duplicatas indesejadas.

 * O uso de IDs de mensagem exclusivos e o cuidado ao lidar com retransmissões são práticas comuns para garantir a indepotência. O Kafka fornece um identificador exclusivo para cada mensagem, e o produtor pode ser configurado para garantir a indepotência.