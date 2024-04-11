# Explicação do código

## CalculateAveragePolar:
-  Leitura do arquivo de dados:
        Utiliza a função scan_csv() do Polars para ler o arquivo CSV.
        Especifica que o arquivo está separado por ponto e vírgula (separator=";").
        Indica que o arquivo não possui cabeçalho (has_header=False).
        Define os nomes das colunas como "station_name" e "measurement" usando uma função lambda (with_column_names=lambda cols: ["station_name", "measurement"]).

- Agrupamento dos dados:
        Agrupa os dados pelo nome da estação usando group_by("station_name").
        Calcula estatísticas de resumo para cada estação, incluindo mínimo, média e máximo das medições de temperatura.
        Renomeia as colunas resultantes usando alias() para "min_measurement", "mean_measurement" e "max_measurement".
        Ordena os dados agrupados por nome de estação usando sort("station_name").
        Coleta os resultados, indicando que a coleta será feita de forma contínua (streaming=True).

- Impressão dos resultados finais:
        Itera sobre as linhas dos resultados agrupados.
        Imprime os resultados formatados no formato nome_da_estação=valor_mínimo/valor_médio/valor_máximo, com uma precisão de uma casa decimal.
        Os resultados são impressos entre chaves {}.

## CreateMeasurements:
- Definição da classe CreateMeasurement:
    Esta classe é responsável por gerar os dados das medições.
    Atributos da classe:
        STATIONS: É uma lista de tuplas contendo o nome da estação e a temperatura média.
        stations: Um DataFrame Polars criado a partir das estações e temperaturas médias.

 -  Método generate_batch:
    Este método gera um lote de dados de medição. Ele amostra aleatoriamente a partir do DataFrame stations, adiciona ruído gaussiano à temperatura e retorna um DataFrame Polars contendo os dados do lote.

 - Método generate_measurement_file:
    Este método gera o arquivo de mediçõesD. Ele recebe parâmetros como nome do arquivo, número de registros a serem gerados, separador e desvio padrão para adicionar ruído à temperatura. Ele divide a geração de dados em lotes e escreve cada lote no arquivo conforme ele é gerado.

- Seção if __name__ == "__main__":
    Aqui, a execução do programa é iniciada. Ele define um analisador de argumentos de linha de comando usando o módulo argparse. Os argumentos possíveis incluem o nome do arquivo de saída e o número de registros a serem gerados. Em seguida, ele instancia a classe CreateMeasurement e chama o método generate_measurement_file com base nos argumentos fornecidos na linha de comando.

-    Função min_records:
    Esta função é usada para validar o argumento records fornecido na linha de comando, garantindo que seja um número inteiro válido e que seja maior ou igual a 1.