#!/bin/sh

# Change the contents of this output to get the environment variables
# of interest. The output must be valid JSON, with strings for both
# keys and values.
cat <<EOF
{
  "ZINC_FIRST_ADMIN_USER": "$ZINC_FIRST_ADMIN_USER",
  "ZINC_FIRST_ADMIN_PASSWORD": "$ZINC_FIRST_ADMIN_PASSWORD",
  "API_PORT": "$API_PORT",
  "ZINCSEARCH_USERNAME": "$ZINCSEARCH_USERNAME",
  "ZINCSEARCH_PASSWORD": "$ZINCSEARCH_USERNAME",
  "ZINCSEARCH_HOST": "$ZINCSEARCH_HOST"
}
EOF
