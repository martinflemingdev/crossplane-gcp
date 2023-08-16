import argparse
import apache_beam as beam
from apache_beam.options.pipeline_options import PipelineOptions

def add_processed(element):
    return element + 'processed'

def run(argv=None):
    parser = argparse.ArgumentParser()
    parser.add_argument('--inputFile', dest='inputFile', required=True)
    parser.add_argument('--outputFile', dest='outputFile', required=True)
    known_args, pipeline_args = parser.parse_known_args(argv)

    pipeline_options = PipelineOptions(pipeline_args)

    with beam.Pipeline(options=pipeline_options) as p:
        (p
         | 'Read from GCS' >> beam.io.ReadFromText(known_args.inputFile)
         | 'Add processed' >> beam.Map(add_processed)
         | 'Write to GCS' >> beam.io.WriteToText(known_args.outputFile))

if __name__ == '__main__':
    run()
