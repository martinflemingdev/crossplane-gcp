# from https://medium.com/swlh/apache-beam-google-cloud-dataflow-and-creating-custom-templates-using-python-c666c151b4bc

import apache_beam as beam
from apache_beam.options.pipeline_options import PipelineOptions

with beam.Pipeline(options=PipelineOptions()) as p:
    lines = p | 'Creating PCollection' >> beam.Create(['Hello', 'Hello Good Morning', 'GoodBye'])
    counts = (
        lines
        | 'Tokenizing' >> (beam.FlatMap(lambda x: x.split(' ')))
        | 'Pairing With One' >> beam.Map(lambda x: (x, 1))
        | 'GroupbyKey And Sum' >> beam.CombinePerKey(sum)
        | 'Printing' >> beam.ParDo(lambda x: print(x[0], x[1]))
        )
