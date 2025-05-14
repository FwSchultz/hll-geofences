import yaml
import uuid

# Function to parse fence coordinates (e.g., A3-A8, C1-C10, A2)
def parse_fence(fence_str):
    coordinates = []
    for part in fence_str.split(', '):
        if '-' in part:
            # Handle range like A3-A8 or C1-C10
            start, end = part.split('-')
            col = start[0]
            start_y = int(start[1:])
            end_y = int(end[1:])
            # Ensure inclusive range
            for y in range(start_y, end_y + 1):
                coordinates.append((col, y))
        else:
            # Handle single coordinate like A2
            col = part[0]
            y = int(part[1:])
            coordinates.append((col, y))
    # Sort by column (A-J) then by Y value to ensure consistent order
    return sorted(coordinates, key=lambda x: (x[0], x[1]))

# Function to generate YAML structure for a fence
def generate_fence_entries(coordinates, map_names, fence_type):
    entries = []
    for col, y in coordinates:
        entry = {
            'X': col,
            '"Y"': y,
            'Condition': {
                'Equals': {
                    'game_mode': ['Warfare'],
                    'map_name': list(map_names)  # Ensure full list is copied
                },
                'LessThan': {
                    'player_count': 50
                }
            }
        }
        entries.append(entry)
    return entries

# Function to process a single YAML file and generate the output
def process_file(input_data, output_filename):
    map_names = input_data['map_name']
    allies_fence = parse_fence(input_data['AlliesFence'])
    axis_fence = parse_fence(input_data['AxisFence'])
    
    output_data = {
        'AxisFence': generate_fence_entries(axis_fence, map_names, 'AxisFence'),
        'AlliesFence': generate_fence_entries(allies_fence, map_names, 'AlliesFence')
    }
    
    # Custom representer to ensure proper indentation and formatting
    def represent_dict(dumper, data):
        return dumper.represent_mapping('tag:yaml.org,2002:map', data, flow_style=False)
    
    yaml.add_representer(dict, represent_dict)
    
    with open(output_filename, 'w', encoding='utf-8') as f:
        yaml.dump(output_data, f, default_flow_style=False, sort_keys=False, allow_unicode=True, indent=2, width=1000)

# Input data from the provided YAML files
input_files = [
    {
        'filename': '3capsonly_alliedToAxisHorizontalMaps.yml',
        'data': {
            'map_name': ['CARENTAN', 'HILL 400', 'HÜRTGEN FOREST', 'MORTAIN'],
            'AlliesFence': 'A3-A8, B3-B8, C3-C8, D3-D8, E3-E8, F3-F8, G3-G8, H3-H8, A2, B2, C2, D2, E2, F2, G2, H2, I2, J2, A9, B9, C9, D9, E9, F9, G9, H9, I9, J9',
            'AxisFence': 'J3-J8, I3-I8, H3-H8, G3-G8, F3-F8, E3-E8, D3-D8, C3-C8, A2, B2, C2, D2, E2, F2, G2, H2, I2, J2, A9, B9, C9, D9, E9, F9, G9, H9, I9, J9'
        }
    },
    {
        'filename': '3capsonly_axisToAlliedHorizontalMaps.yml',
        'data': {
            'map_name': ['EL ALAMEIN', 'OMAHA BEACH', 'SAINTE-MÈRE-ÉGLISE', 'STALINGRAD', 'TOBRUK', 'UTAH BEACH'],
            'AlliesFence': 'J3-J8, I3-I8, H3-H8, G3-G8, F3-F8, E3-E8, D3-D8, C3-C8, A2, B2, C2, D2, E2, F2, G2, H2, I2, J2, A9, B9, C9, D9, E9, F9, G9, H9, I9, J9',
            'AxisFence': 'A3-A8, B3-B8, C3-C8, D3-D8, E3-E8, F3-F8, G3-G8, H3-H8, A2, B2, C2, D2, E2, F2, G2, H2, I2, J2, A9, B9, C9, D9, E9, F9, G9, H9, I9, J9'
        }
    },
    {
        'filename': '3capsonly_alliedToAxisVerticalMaps.yml',
        'data': {
            'map_name': ['ELSENBORN RIDGE', 'KHARKOV', 'KURSK', 'PURPLE HEART LANE', 'ST MARIE DU MONT'],
            'AlliesFence': 'C1-C8, D1-D8, E1-E8, F1-F8, G1-G8, H1-H8, B1-B10, I1-I10',
            'AxisFence': 'C3-C10, D3-D10, E3-E10, F3-F10, G3-G10, H3-H10, B1-B10, I1-I10'
        }
    },
    {
        'filename': '3capsonly_axisToAlliedVerticalMaps.yml',
        'data': {
            'map_name': ['DRIEL', 'FOY', 'REMAGEN'],
            'AlliesFence': 'C3-C10, D3-D10, E3-E10, F3-F10, G3-G10, H3-H10, B1-B10, I1-I10',
            'AxisFence': 'C1-C8, D1-D8, E1-E8, F1-F8, G1-G8, H1-H8, B1-B10, I1-I10'
        }
    }
]

# Process each input file
for input_file in input_files:
    output_filename = f"generated_{input_file['filename']}"
    process_file(input_file['data'], output_filename)
    print(f"Generated {output_filename}")
