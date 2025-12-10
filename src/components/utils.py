import os
import sys
import logging
import json

logger = logging.getLogger(__name__)

def get_project_root() -> str:
    return os.path.dirname(os.path.dirname(os.path.abspath(__file__)))

def get_config_file_path(file_name: str) -> str:
    project_root = get_project_root()
    return os.path.join(project_root, 'config', file_name)

def load_config(file_name: str) -> dict:
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

def setup_logging(log_level: str = 'INFO', log_file: str = None) -> None:
    log_format = '%(asctime)s [%(levelname)s] %(message)s'
    log_level = getattr(logging, log_level.upper())
    
    if log_file:
        handlers = [
            logging.FileHandler(log_file),
            logging.StreamHandler()
        ]
    else:
        handlers = [logging.StreamHandler()]
    
    logging.basicConfig(
        format=log_format,
        level=log_level,
        handlers=handlers
    )