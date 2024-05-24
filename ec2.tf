# Create Security Group - Web Traffic
resource "aws_security_group" "vpc-email-searcher" {
  name        = "vpc-email-searcher"
  description = "Allow SSH and HTTP access"

  ingress {
    from_port   = 22
    to_port     = 22
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Allow SSH from anywhere
  }

  ingress {
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"] # Allow HTTP from anywhere
  }

  ingress {
    from_port   = 3000
    to_port     = 3000
    protocol    = "tcp"
    cidr_blocks = ["0.0.0.0/0"]
  }

  egress {
    description = "Allow all ip and ports outbound"
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

# Resource: EC2 Instance
resource "aws_instance" "email-searcher" {
  ami                    = data.aws_ami.amzLinux.id
  instance_type          = "t2.micro"
  vpc_security_group_ids = [aws_security_group.vpc-email-searcher.id]
  user_data = templatefile("initialize-app-script.sh", {
    ZINC_FIRST_ADMIN_USER     = data.external.env.result["ZINC_FIRST_ADMIN_USER"],
    ZINC_FIRST_ADMIN_PASSWORD = data.external.env.result["ZINC_FIRST_ADMIN_PASSWORD"],
    API_PORT                  = data.external.env.result["API_PORT"],
    ZINCSEARCH_USERNAME       = data.external.env.result["ZINCSEARCH_USERNAME"],
    ZINCSEARCH_PASSWORD       = data.external.env.result["ZINCSEARCH_PASSWORD"],
    ZINCSEARCH_HOST           = data.external.env.result["ZINCSEARCH_HOST"]
    FRONT_END_ADD             = "http://${aws_instance.email-searcher.public_dns}:8080"
  })

  tags = {
    Name = "EmailSearcherApp"
  }

}
