
# Steps to Create the Proof of Concept and Demo

1. Create the Mock Scripts and go-example Program
    - Setup Script: Write a mock shell script (setup-env.sh) that simply prints "Setting up environment...".
    - Teardown Script: Write a mock shell script (teardown-env.sh) that prints "Tearing down environment...".
    - Go Program:
        - Create the go-example program that:
            - Takes an optional -p argument for the percentage chance of passing.
            - Takes an optional -f argument for reading a text file.
            - Prints the contents of the file if provided, or prints "no file received".
            - Outputs go-example-fail.txt if it fails, containing the failure message and percentage.
1. Set Up the GitHub Repository
    - Create a new GitHub repository.
    - Add the mock scripts and the Go program to the repository.
    - Add a .github/workflows directory to store the GitHub Actions workflow YAML file.
1. Write the GitHub Actions Workflow
    - Define a workflow YAML file to:
        - take an input maxRetry with a default value (e.g., 3)
        - create three separate jobs for setup, execution, and teardown
        - upload the go-example-fail.txt artifact if the go-example job fails
        - re-run the workflow up to maxRetry times if the execution job fails
1. Test the Workflow Locally
    - Run the setup script, Go program, and teardown script locally to ensure they behave as expected.
    - Test different percentages and file inputs for go-example.
1. Push the Code to GitHub
    - Push the repository to GitHub to trigger the workflow.
    - Check the behavior of the workflow, ensuring retries and artifact handling work as expected.
1. Iterate and Improve
    - Analyze the workflow execution logs and refine the workflow if needed.
    - Adjust the mock scripts or Go program to better mimic the real-world scenario.


# Questions to Help Flesh Out the Details
Regarding the Mock Scripts and Go Program:
- What language should the setup and teardown scripts be written in? (e.g., Bash, Python)
- Should the failure message in the artifact be customizable or fixed?
- Do you want specific logging/formatting for the go-example outputs, or is plain text sufficient?
Regarding the Workflow:
- What is the default value for maxRetry in the workflow? Should it be user-configurable at the workflow level?
- Should retries happen immediately, or should there be a delay between retries? If so, how long?
- Should the workflow fail entirely if it exhausts retries, or should it provide additional outputs (e.g., summary logs)?
Regarding Artifacts:
- Is uploading artifacts the only way to pass the failure file to subsequent runs, or do you prefer a different method?
- Should all failure artifacts from retries be saved, or just the latest one?
Regarding Workflow Execution:
- Should the workflow run in parallel with other builds, or is it strictly sequential?
- How should job dependencies be handled to ensure the teardown always runs after execution, even if execution fails?
