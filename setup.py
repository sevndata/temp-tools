from setuptools import setup, find_packages

setup(
    name='temp-tools',
    version='0.1',
    packages=find_packages(),
    install_requires=[
        'requests'
    ],
    tests_require=[
        'pytest',
    ],
    python_requires='>=3.0',
)
