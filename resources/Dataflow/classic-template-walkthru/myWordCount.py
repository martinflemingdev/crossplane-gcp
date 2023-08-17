

class WordcountOptions(PipelineOptions):
    @classmethod
    def _add_argparse_args(cls, parser):
        parser.add_argument(
            '--input',
            default='gs://dataflow-samples/shakespeare/kinglear.txt',
            help='Path of the file to read from')
        parser.add_argument(
            '--output',
            required=True,
            help='Output file to write results to.')
    pipeline_options = PipelineOptions(['--output', 'some/output_path'])
    p = beam.Pipeline(options=pipeline_options)

    wordcount_options = pipeline_options.view_as(WordcountOptions)
    lines = p | 'read' >> ReadFromText(wordcount_options.input)