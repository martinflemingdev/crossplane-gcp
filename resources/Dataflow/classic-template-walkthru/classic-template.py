import argparse
import apache_beam as beam
from apache_beam.options.pipeline_options import PipelineOptions

def run(argv=None):
    parser = argparse.ArgumentParser()
    parser.add_argument('--inputFile', required=True, help='GCS path to the input file.')
    parser.add_argument('--outputFile', required=True, help='GCS path for the output file.')
    known_args, pipeline_args = parser.parse_known_args(argv)

    pipeline_options = PipelineOptions(pipeline_args)

    with beam.Pipeline(options=pipeline_options) as p:
        (
            p | 'Read from GCS' >> beam.io.ReadFromText(known_args.inputFile)
              | 'Write to GCS' >> beam.io.WriteToText(known_args.outputFile)
        )

if __name__ == '__main__':
    run()