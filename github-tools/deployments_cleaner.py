import requests
import time

# Replace with your GitHub username
username = ''

# Replace with your GitHub token, [personal access token](https://github.com/settings/tokens)
token = ''

# Replace with your GitHub repository name
repository_name = ''

base_url = f'https://api.github.com/repos/{username}/{repository_name}'
deployments_url = f'{base_url}/deployments'
statuses_url = f'{base_url}/deployments/{{}}/statuses'


# Function to get deployments
def get_deployments():
    response = requests.get(deployments_url, auth=(username, token))
    if response.status_code == 200:
        return response.json()
    else:
        response.raise_for_status()


# Function to mark a deployment as inactive
def inactive_deployment(deployment_id):
    url = statuses_url.format(deployment_id)
    data = {"state": "inactive"}
    response = requests.post(url, json=data, auth=(username, token),
                             headers={"Content-Type": "application/json",
                                      "Accept": "application/vnd.github.ant-man-preview+json"})
    if response.status_code == 201:
        print(f"Deployment {deployment_id} marked as inactive.")
        return True
    else:
        print(f"Failed to update deployment {deployment_id}. Status Code: {response.status_code}")
        return False


# Function to delete a deployment
def delete_deployment(deployment_id):
    if inactive_deployment(deployment_id):
        url = f'{deployments_url}/{deployment_id}'
        response = requests.delete(url, auth=(username, token))
        return response
    else:
        return None


# Function to delete all deployments found in this query
def delete_deployments_this_query():
    try:
        deployments = get_deployments()
    except requests.HTTPError as e:
        print(f"Failed to retrieve deployments. Error: {e}")
        return False

    if not deployments:
        return False

    for deployment in deployments:
        deployment_id = deployment.get("id")
        if deployment_id is not None:
            response = delete_deployment(int(deployment_id))
            if response and response.status_code != 204:
                print(f"Failed to delete deployment {deployment_id}. Status Code: {response.status_code}")

    return True


# Function to delete all deployments
def delete_all_deployments():
    while True:
        try:
            if not delete_deployments_this_query():
                break
            time.sleep(2)
        except Exception as e:
            print(f"An error occurred: {e}")
            break
