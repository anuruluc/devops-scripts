import os
import sys
import logging

# Define the logger
logger = logging.getLogger(__name__)

def get_project_root() -> str:
    """Returns project root directory"""
    return os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

def get_config_file_path(file_name: str) -> str:
    """Returns the full path to the given config file"""
    project_root = get_project_root()
    return os.path.join(project_root, 'config', file_name)

def load_config(file_name: str) -> dict:
    """Loads the given config file and returns its contents as a dictionary"""
    import json
    file_path = get_config_file_path(file_name)
    try:
        with open(file_path, 'r') as file:
            return json.load(file)
    except FileNotFoundError:
        logger.error(f"Config file '{file_name}' not found")
        sys.exit(1)
    except json.JSONDecodeError as e:
        logger.error(f"Failed to parse config file '{file_name}': {e}")
        sys.exit(1)

def setup_logging(log_level: str = 'INFO') -> None:
    """Sets up the logging configuration"""
    logging.basicConfig(
        format='%(asctime)s [%(levelname)s] %(message)s',
        level=getattr(logging, log_level.upper())
    )