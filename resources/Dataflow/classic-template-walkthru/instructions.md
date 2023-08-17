https://cloud.google.com/dataflow/docs/guides/templates/creating-templates#python

In this document, you learn how to create a custom **classic template** from your Dataflow pipeline code. 

Classic templates package existing Dataflow pipelines to create reusable templates that you can customize for each job by changing specific pipeline parameters. 

**Rather than writing the template, you use a command to generate the template from an existing pipeline.**

The following is a brief overview of the process. Details of this process are provided in subsequent sections.

1. In your pipeline code, use the ValueProvider interface for all pipeline options that you want to set or use at runtime. Use DoFn objects that accept runtime parameters.
2. Extend your template with additional metadata so that custom parameters are validated when the classic template is run. Examples of such metadata include the name of your custom classic template and optional parameters.
3. Check if the pipeline I/O connectors support ValueProvider objects, and make changes as required.
4. Create and stage the custom classic template.
5. Run the custom classic template.

**The Dataflow runner does not support ValueProvider options for Pub/Sub topics and subscription parameters. If you require Pub/Sub options in your runtime parameters, use Flex Templates.**

