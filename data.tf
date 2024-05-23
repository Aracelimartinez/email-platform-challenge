# Get the ami
data "aws_ami" "amzLinux" {
  most_recent = true
  owners      = ["137112412989"]

  filter {
    name   = "name"
    values = ["al2023-ami-2023.4.20240513.0-kernel-6.1-x86_64"]
  }

}

# Run the script to get the environment variables to run the docker-compose file.
data "external" "env" {
  program = ["${path.module}/initialize-env-script.sh"]
}
